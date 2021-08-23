# Blackmagic Camera Control Protocol

## Now in a machine-readable format

The Blackmagic line of cameras are extremely powerful for the price. Especially considering their ability to be controlled both over SDI and Bluetooth, making this a very fun camera to use in situations where you need absolute control.

This repo contains a machine-readable file called `PROTOCOL.json` which you can use to build your own application around the camera. `PROTOCOL.json` contains both the documentation of the parameters from both PDFs (camera protocol & pocket cameras) + information about the bluetooth services. This is a small example of how the file is structured:

```json
{
	"information": "",
	"groups": [
		{
			"name": "Lens",
			"normalized_name": "lens",
			"id": 0,
			"parameters": [
				{
					"id": 0,
					"group": "Lens",
					"group_id": 0,
					"parameter": "Focus",
					"type": "fixed16",
					"index": [],
					"minimum": 0,
					"maximum": 1,
					"interpretation": "0.0 = near, 1.0 = far"
				}
			]
		}
	],
	"bluetooth_services": [
		{
			"name": "Blackmagic Camera Service",
			"normalized_name": "blackmagic_camera_service",
			"uuid": "291d567a-6d75-11e6-8b77-86f30ca893d3",
			"characteristics": [
				{
					"name": "Outgoing Camera Control",
					"normalized_name": "outgoing_camera_control",
					"uuid": "5dd3465f-1aee-4299-8493-d2eca2f8e1bb",
					"description": "Send Camera Control messages"
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
