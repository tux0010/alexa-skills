package main

import (
	"fmt"
	"strings"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
	log "github.com/sirupsen/logrus"
)

// TODO: Use go generate to build this
var cuisines = []string{"italian", "french", "thai", "japanese", "indian", "korean", "chinese", "vietnamese", "mexican", "american"}

var Application = map[string]interface{}{
	"/echo/restaurant_finder": alexa.EchoApplication{
		AppID: "amzn1.ask.skill.4fa65756-1e28-4779-9a1a-28e175a9c609",
		//Handler: EchoRestaurantFinder,
		OnIntent: EchoRestaurantFinder,
	},
}

func EchoRestaurantFinder(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	switch echoReq.GetRequestType() {
	case "LaunchRequest":
		msg := "Hello. You can ask me for a restaurant recommendation, based on a cuisine."
		msg += " To start, ask me about the available cuisines. Then you can ask me a question"
		msg += " like, find me an Italian restaurant."
		alexa.NewEchoResponse().OutputSpeech(msg).EndSession(false)
	case "IntentRequest":
		log.WithFields(log.Fields{
			"request_type": "IntentRequest",
			"intent_name":  echoReq.GetIntentName(),
			"slots":        echoReq.AllSlots(),
		}).Info("got Echo request")

		switch echoReq.GetIntentName() {
		case "Question":
			cuisine, err := echoReq.GetSlotValue("Cuisine")
			if err != nil {
				log.Error(err)
				return
			}

			cuisine = strings.ToLower(cuisine)
			r, err := restaurantRecommendation(cuisine)
			if err != nil {
				log.Error(err)
				return
			}

			msg := fmt.Sprintf("I recommend %s, located at %s. The details have been added to the Alexa app", r.Name, r.Address)
			card := strings.Join([]string{r.Name, r.Address, r.Phone}, "\n")
			echoResp.OutputSpeech(msg).Card("Question", card).EndSession(true)

		default:
			msg := "You can ask me for a restaurant recommendation, based on a cuisine type."
			msg += " You can also ask me to list all available cuisine types."
			alexa.NewEchoResponse().OutputSpeech(msg).EndSession(false)
		}
	case "SessionEndedRequest":
		alexa.NewEchoResponse().OutputSpeech("Bye").EndSession(true)
	}
}

func main() {
	alexa.Run(Application, "3000")
}
