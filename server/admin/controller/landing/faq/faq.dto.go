package faq

type FaqDto struct {
	ID    	    string 		`param:"id"`
	Name 		string 		`json:"name"`
	Answer 		string 		`json:"answer"`
	Question  	string 		`json:"question"`
}
