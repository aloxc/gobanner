FROM scratch
LABEL maintainer="leerohwa@gmail.com"
LABEL add="https://github.com/aloxc/gobanner"
LABEL version="1.0"
LABEL description="this is a golang banner,using golang file create a banner"

RUN wget "https://github.com/aloxc/gobanner/gobanner"
RUN chmod +x gobanner
RUN gobanner