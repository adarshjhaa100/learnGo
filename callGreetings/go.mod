module learningo/callGreetings

go 1.19

replace learningo/greetings => ../greetings

replace learningo/randomgreetings => ../randomGreetings

require (
	learningo/greetings v0.0.0
	learningo/randomgreetings v0.0.0-00010101000000-000000000000
)
