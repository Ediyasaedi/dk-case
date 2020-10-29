# dk-case

1. In terminal, go to this file directory
2. running command: docker build -t <name_of_new_docker_images> <path_Dockerfile> <br />
   example: `docker build -t dk-case-images . `
3. running command: docker run -p <your_local_port>:8000 <name_of_new_docker_images> <br />
   example: `docker run -p 8080:8000 dk-case-images`
4. Test all routes at api-doc.md with your local port using Postman or any application for rest-api test
