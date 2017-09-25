FROM debian:wheezy
MAINTAINER Sundara Senthil

ADD .build/server /bin/

ADD run.sh /run.sh

RUN chmod 755 /run.sh
RUN chmod -R 755 /bin/*

CMD ["/run.sh"]
