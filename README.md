
# goForward

## Description  
Simple port fowarding app 

### Docker Build
The dockerfile in this project uses a multi stage docker build which requires docker >= 17.05 


### Example  
CLI  
```
goforward -l 8080 -dh google.com -dp 80
```

Docker
```
docker run -ti -e LISTEN_PORT=8080 -e DEST_HOST=google.com -e DEST_PORT=80 -p 8080:8080 gofoward:latest
```

## Documentation
Simple tcp port forwarding app that listens on a port, then fowards all TCP traffic to the specified destination host and port. 
### options
Available command line optoins  

```
   --listen value, -l value       Posrt to listen on [$LISTEN_PORT]
   --dest-host value, --dh value  Host to forward traffic to [$DEST_HOST]
   --dest-port value, --dp value  Port to forward traffit to [$DEST_PORT]
   --help, -h                     show help
   --version, -v                  print the version
```

If no options are passed the app will fallback to looking at ENV vars. This is done for easy running as a container. 

## Roadmap
The next planned features are UDP port fowarding and ability to forward multiple ports to multiple destinations. 