FROM scratch

WORKDIR /Users/jim/workspace/go-blog-step-by-step
COPY . /Users/jim/workspace/go-blog-step-by-step

EXPOSE 8000
CMD ["./go-blog-step-by-step"]