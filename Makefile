test:
	curl -XPOST https://projectname-227718.appspot.com -d @curl.json

deploy:
	gcloud app deploy