all: hello skilstak README.html

clean:
	if [ -e index.html ]; then \
		rm index.html;\
		rm *html;\
	fi

%.html: %.md
	pandoc $^ -o $@

index.html:
	curl https://skilstak.io -o index.html

skilstak: index.html
