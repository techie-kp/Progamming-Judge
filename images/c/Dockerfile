FROM frolvlad/alpine-gcc

# Create the dummy user
RUN adduser -D -s /bin/sh execution_user
# Copy the execution command
COPY evaluate.sh /home/execution_user/
RUN chmod +x /home/execution_user/evaluate.sh

# Dummy user to perform the execution
USER execution_user
WORKDIR /home/execution_user
ENTRYPOINT [ "./evaluate.sh" ]     