FROM scratch
# Copy our static executable.
COPY ./bank-account /bank-account
# Run the hello binary.
ENTRYPOINT ["/bank-account"]