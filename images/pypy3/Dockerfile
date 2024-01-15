FROM pypy:3.8-7.3.7-slim-bullseye

# Create the dummy user
RUN useradd --create-home --shell /bin/bash execution_user

# Copy the execution command
COPY evaluate.sh /home/execution_user/
RUN chmod +x /home/execution_user/evaluate.sh

# Dummy user to perform the execution
USER execution_user
WORKDIR /home/execution_user
ENTRYPOINT [ "./evaluate.sh" ]