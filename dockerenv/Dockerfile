FROM debian:wheezy
MAINTAINER Sundara Senthil <senthil.sundara@gmail.com>

ADD .build/dockernenv /bin/
ADD run.sh /run.sh

RUN chmod 755 /run.sh
RUN chmod -R 755 /bin/*

CMD ["/run.sh"]

