FROM alpine:3.8

COPY ./build /runnable
ENV PATH='$PATH:/runnable'

CMD ["mailer"]

