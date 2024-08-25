// Package controller provides utilities for managing route handlers in an Echo framework application.
// It enhances Echo's capabilities by adding additional features to route handlers and context management.
// This package facilitates the creation of robust web applications by extending Echo's functionality.
//
// Features:
//   - Enhanced Context Management: The Context type wraps Echo's standard context (echo.Context)
//     and adds additional functionalities for handling web requests.
//   - Middleware Initialization: The Initialize function creates middleware for initializing a custom Context instance,
//     allowing extended features to be added to route handlers.
//   - Route Handler Registration: The Register function registers route handlers with additional features
//     and replaces Echo's context with the custom controller.Context type.
//   - HTML Rendering: The Html method renders templ components and returns them as HTML,
//     supporting both full page rendering and fragment rendering for htmx requests.
//   - htmx Request Detection: The IsHtmx method checks if a request is made via htmx by examining the "Hx-Request" header.
package controller

import (
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"main/internal/model"
	"main/pkg/globals"
	admin "main/server/admin/view/pages"
	index "main/server/app/v1/view/pages"
)

// Context wraps the Echo's standard context (echo.Context) to provide additional functionalities.
//
// Functions added by it:
//
// - Initialize() echo.MiddlewareFunc
//
// - Register() echo.HandlerFunc
//
// - Html() error
//
// - IsHtmx() bool
type Context struct {
	echo.Context
}

// Initialize creates a middleware function for initializing a controller Context instance.
// It wraps the standard Echo context with the controller Context type, allowing additional functionalities to be added.
//
// Example usage:
//
//	app := echo.New()
//	app.Use(controller.Initialize())
//
// Returns:
//   - A middleware function that initializes a controller Context instance.
//
// Notes:
//   - This function is intended to be used as middleware in an Echo application.
//   - It wraps the standard Echo context with the controller Context type, allowing additional functionalities to be added.
//   - The controller Context instance can be accessed within route handlers to utilize the extended features provided.
func Initialize() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &Context{Context: c}
			return next(ctx)
		}
	}
}

// Registers echo's route handlers with additional features to an Echo's Context instance.
// It takes a route handler function as an argument, enhances it with additional functionality,
// and replaces the Echo context with the custom controller.Context type to provide extended features.
//
// Example usage:
//
//	  app := echo.New()
//	  app.GET("/path", controller.Register(func(ctx *controller.Context) error {
//			if ctx.IsHtmx() {
//				return ctx.Html(view.Fragment())
//	     }
//
//	     return ctx.Html(view.FullPage())
//	  }))
//
// Parameters:
//   - app: An instance of Echo framework where the route handler will be registered.
//   - func(ctx *controller.Context): The route handler function to be registered. It should have the signature func(Context) error.
//
// Notes:
//   - The route handler function should take controller.Context as an argument to utilize the additional features provided by this package.
//   - The controller.Context type extends the standard echo.Context with extra methods and features, like Html() IsHtmx() and others.
func Use(handlerFunc func(*Context) error) echo.HandlerFunc {
	return func(c echo.Context) error { return handlerFunc(c.(*Context)) }
}

func Set[T any](handlerFunc interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		DTOstruct := new(T)
		if err := c.Bind(DTOstruct); err != nil {
			return c.String(http.StatusBadRequest, "Parameters Binding Problem: "+err.Error())
		}

		if reflect.TypeOf(DTOstruct).String() != "*interface {}" {
			if err := Validate(DTOstruct); err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
		}

		arguments := []reflect.Value{reflect.ValueOf(c.(*Context))}
		if reflect.TypeOf(DTOstruct).String() != "*interface {}" {
			argValue := reflect.New(reflect.TypeOf(DTOstruct)).Elem()
			argValue.Set(reflect.ValueOf(DTOstruct))
			arguments = append(arguments, argValue)
		}

		result := reflect.ValueOf(handlerFunc).Call(arguments)
		if result[0].IsNil() {
			return nil
		}
		return result[0].Interface().(error)
	}
}

func Validate(DTOstruct interface{}) error {
	var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(DTOstruct)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		// _, err := err.(validator.ValidationErrors)
		// for _, err := range err.(validator.ValidationErrors) {
		// 	errors += err.Error() + "  |  "
		// }

		// fmt.Println(err.Namespace())
		// fmt.Println(err.Field())
		// fmt.Println(err.StructNamespace())
		// fmt.Println(err.StructField())
		// fmt.Println(err.Tag())
		// fmt.Println(err.ActualTag())
		// fmt.Println(err.Kind())
		// fmt.Println(err.Type())
		// fmt.Println(err.Value())
		// fmt.Println(err.Param())
		// fmt.Println()
		// from here you can create your own error messages in whatever language you wish
		return err
	}
	return nil
}

// Html renders the given templ component and returns it as HTML.
// If the request is made via htmx, it returns the component as a fragment.
// Otherwise, it embeds the component within the layout of the base HTML.
//
// Example usage:
//
//	app := echo.New()
//	app.GET("/fullpage", controller.Register(func(ctx *controller.Context) error {
//	   return ctx.Html(view.FullPage())
//	}))
//
// Parameters:
//   - component: The templ component to be rendered.
//
// Returns:
//   - An error, if any, encountered during rendering.
//
// Notes:
//   - The method requires access to the request context and response writer via the Context instance.
//   - If the request is made via htmx, the component is rendered as a fragment without any layout.
//   - If the request is not made via htmx, the component is rendered within the layout of the base HTML.
//   - Make sure to handle any errors returned by this method appropriately.
func (ctx *Context) Html(component templ.Component) error {
	return ctx.HtmlWithStatus(http.StatusOK, component)
}

func (ctx *Context) HtmlWithStatus(code int, component templ.Component) error {
	if ctx.IsHtmx() {
		return component.Render(ctx.Request().Context(), ctx.Response())
	}

	var Base templ.Component
	if !ctx.IsAdmin() {
		Base = index.Layout(ctx.Get("Interface").(model.Interface), component)
	} else {
		Base = admin.Page(component)
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	ctx.Response().Writer.WriteHeader(code)
	return Base.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (ctx *Context) Renders(code int, component templ.Component) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	ctx.Response().Writer.WriteHeader(code)
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (ctx *Context) RenderPlain(component templ.Component) string {
	// component.Render(context.Background(), os.)
	return ""
}

func (ctx *Context) IsAdmin() bool {
	if ctx.Get("ISADMIN") == nil {
		return false
	}
	return ctx.Get("ISADMIN").(bool)
}

func (ctx *Context) User() model.Users {
	return ctx.Get("USER").(model.Users)
}

// IsHtmx checks if the request is made via htmx (Hypertext Markup eXtension).
// It examines the request headers and returns true if the "Hx-Request" header is set to "true".
//
// Example usage:
//
//	isHtmxRequest := ctx.IsHtmx()
//
// Returns:
//   - true if the request is made via htmx, false otherwise.
//
// Notes:
//   - The method requires access to the request context via the Context instance.
//   - It examines the "Hx-Request" header to determine if the request is an htmx request.
//   - This method can be used to conditionally render content or handle logic based on the type of request.
func (ctx *Context) IsHtmx() bool {
	return ctx.Request().Header.Get("Hx-Request") == "true" && ctx.Request().Header.Get("hx-fullPage") != "true"
}

type Cookie struct {
	Key     string
	Value   string
	Expires time.Time
}

func (ctx *Context) WriteCookie(data Cookie) {
	cookie := new(http.Cookie)
	cookie.Name = data.Key
	cookie.Value = data.Value
	cookie.Expires = data.Expires
	ctx.SetCookie(cookie)
}

func (ctx *Context) ReadCookie(Key string) Cookie {
	cookie, err := ctx.Cookie(Key)
	if err != nil {
		cookie = &http.Cookie{Name: "", Value: "", Expires: time.Now()}
	}
	return Cookie{Key: cookie.Name, Value: cookie.Value, Expires: cookie.Expires}
}

type QueryPageParameter struct {
	Page     string `query:"page"`
	PageSize string `query:"pageSize"`
}

func (ctx *Context) Page() int {
	var Query QueryPageParameter

	if ctx.QueryParam("page") == "" {
		Query.Page = "1"
	} else {
		ctx.Bind(&Query)
	}

	page, _ := strconv.Atoi(Query.Page)
	if page <= 0 {
		page = 1
	}
	return page
}

func (ctx *Context) PageSize() int {
	var Query QueryPageParameter
	if ctx.QueryParam("pageSize") == "" {
		Query.PageSize = "-1"
	}

	ctx.Bind(&Query)
	pageSize, _ := strconv.Atoi(Query.PageSize)

	if pageSize <= 0 || pageSize > globals.Env.PageMaxSize {
		pageSize = globals.Env.PageMaxSize
	}
	return pageSize
}
