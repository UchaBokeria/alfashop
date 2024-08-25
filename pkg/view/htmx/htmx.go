package htmx

import "github.com/a-h/templ"

func Link(path string) templ.Attributes {
	return templ.Attributes{
		"hx-get":       path,
		"hx-swap":      "innerHTML show:window:top",
		"hx-push-url":  "true",
		"hx-target":    "#Content",
		"hx-encoding":  "text/html",
		"hx-indicator": ".Loading",
		"hx-trigger":   "click",
		// "hx-boost":                    "true",
		// "hx-push":                     "true",
		// "hx-boost-children":           "true",
		// "hx-boost-ignore":             "true",
		// "hx-boost-swap":               "true",
		// "hx-boost-children-swap":      "true",
		// "hx-boost-ignore-swap":        "true",
		// "hx-boost-swap-mode":          "replace",
		// "hx-boost-children-swap-mode": "replace",
		// "hx-boost-ignore-swap-mode":   "replace",
	}
}
