#Base Image
FROM golang



#WORKDIR
WORKDIR /apps


#Copy All Code
COPY . .

#Build App
RUN go build -o apps main.go


#RUN APP
CMD ["./apps"]



