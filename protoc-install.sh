    # Make sure you grab the latest version
    curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
    # Unzip
    unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
    # Move protoc to /usr/local/bin/
    mv bin/* /usr/local/bin/
    # Move protoc3/include to /usr/local/include/
    mv include/* /usr/local/include/
    # Optional: change owner
    sudo chown alvin /usr/local/bin/protoc
    sudo chown -R alvin /usr/local/include/google