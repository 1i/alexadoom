package main

import (
	"fmt"
	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

// Handler is the lambda hander
func Handler() (alexa.Response, error) {
	fmt.Print("start ")

	text := "<speak><amazon:effect name='whispered'> In the first age, in the first battle, when the shadows first lengthened, one stood. He chose the path of perpetual torment. In his ravenous hatred he found no peace. And with boiling blood he scoured the Umbral Plains seeking vengeance against the dark lords who had wronged him. And those that tasted the bite of his sword named him... </amazon:effect> the Doom Slayer. </speak>"
	payload := alexa.Payload{SSML: text, Type: "SSML"}
	var response = alexa.Response{
		Version:           "",
		SessionAttributes: nil,
		Body:              alexa.ResBody{OutputSpeech: &payload, ShouldEndSession: true},
	}

	spew.Dump(response)
	fmt.Sprintf("response : %T/n", response)
	fmt.Print("end ")
	return response, nil
	//return alexa.NewSimpleResponse("Saying Hello", "<speak><amazon:effect name='whispered'>" + "In the first age, in the first battle, when the shadows first lengthened, one stood. He chose the path of perpetual torment. In his ravenous hatred he found no peace. And with boiling blood he scoured the Umbral Plains seeking vengeance against the dark lords who had wronged him. And those that tasted the bite of his sword named him... the Doom Slayer.  " + "</amazon:effect></speak>"), nil
}

func main() {
	lambda.Start(Handler)
}
