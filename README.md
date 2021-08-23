# Blackmagic Camera Control Protocol

## Now in a machine-readable format

The Blackmagic line of cameras are extremely powerful for the price. Especially considering their ability to be controlled both over SDI and Bluetooth, making this a very fun camera to use in situations where you need absolute control.

This repo contains a machine-readable file called `PROTOCOL.json` which you can use to build your own application around the camera. This is a small example of how the file is structured:

```json
{
	"information": "", //General information about the file
	"groups": [
		{
			"name": "Lens", // Group name
			"normalized_name": "lens", //Normalized name (snake case)
			"id": 0, //Group id
			"parameters": [
				{
					"id": 0, //Parameter id
					"group": "Lens", //Group name (easy reverse lookup)
					"group_id": 0, //Group id (easy reverse lookup)
					"parameter": "Focus", //Parameter name
					"type": "fixed16", //Type
					"index": [], //Index map of paginated settings
					"minimum": 0, //Min
					"maximum": 1, //Max
					"interpretation": "0.0 = near, 1.0 = far" //Info from manual
				}
			]
		}
	],
	"bluetooth_services": [
		{
			"name": "Blackmagic Camera Service", //Name of service
			"normalized_name": "blackmagic_camera_service", //Normalized (snake_case)
			"uuid": "291d567a-6d75-11e6-8b77-86f30ca893d3", //UUID of service
			"characteristics": [
				{
					"name": "Outgoing Camera Control", //Name of characteristics
					"normalized_name": "outgoing_camera_control", //Normalized (snake_case)
					"uuid": "5dd3465f-1aee-4299-8493-d2eca2f8e1bb", //UUID of characteristics
					"description": "Send Camera Control messages" //Description from manual
				}
			]
		}
	]
}
```

I also provided some demos of how to deserialize this file (CALL ME IF YOU GET LOST LUL). You can find them under `/usage`. For the Go demo, just `cd` to the folder and run `go run definitions.go`. For the Rust demo, `cd` to the folder and run `cargo run`.

## Prompt to Blackmagic Design

The cameras Blackmagic makes are amazing, there is no doubt about it. The developer experience around them however aren't. While it probably sounded like a terrific idea on paper to just use the same protocol that works over SDI it feels like missing 20 years of progress in software engineering. Why should I have to manually decode some vendor specific protocol when we have tools such as [Protobuf](https://developers.google.com/protocol-buffers)?

On top of that, the documentation provided for this protocol is "truly broadcast" in the sense that it's atrocious. Sure, plus points to the fact that it actually exists in the open in the first place compared to many other vendors but if you actually have the data that generated that table in the PDF, why not supplement this documentation with a machine-readable file? It doesn't have to be JSON, it could be XML for all I care but I just don't find it funny implementing this from a PDF.

## Contributing

Submit a PR LUL
