#!/usr/bin/env python

# Build utterances
utterances = [
    'Question find me a {Cuisine} restaurant',
    'Question find me an {Cuisine} restaurant',
    'Question find me a good {Cuisine} restaurant',
    'Question recommend a {Cuisine} restaurant',
    'Question recommend an {Cuisine} restaurant',
    'Question recommend a good {Cuisine} restaurant',
    'Question what\'s a good {Cuisine} restaurant',
    'Question what\'s a good {Cuisine} restaurant near me',
]

with open('utterances.txt', 'w') as f:
    for u in utterances:
        f.write(u + '\n')

with open('intent_schema.json', 'w') as f:
    schema = '''
{
  "intents": [
    {
      "intent": "Question",
      "slots": [
        {
          "name": "Cuisine",
          "type": "LIST_OF_CUISINES"
        }
      ]
    }
  ]
}
'''

    f.write(schema.strip())

print 'utterances & schema written'