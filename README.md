# wahanthok
A simple online dictionary written in golang for Google App Engine

The app is available at http://wahanthok.appspot.com . You can search for an English word and find its meaning in Manipuri and English.

## Using it for your language
The meanings are stored as maps in a gob. If you want to use for your language, all you have to do is replace the gob file. Rename the app identifier in the app.yaml file to your new app identifier. Then deploy it on GAE or run it locally.

## Making the gob file
Right now there is no concrete tool to do that. I just wrote a simple program to create the gob with a few words. Maybe you can write one and share it ;) . Anyways create it whatever way you want for now.

## Running locally or deploying it on Google App Engine
You need to install Go Google App Engine SDK https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go
All the instructions are there on the Go GAE documentation. Please read from there on how run locally or deploy it on GAE.
