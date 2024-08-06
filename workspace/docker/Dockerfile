FROM ubuntu:latest
LABEL authors="bxcode"

# Update the package repository and install openssh-server
RUN apt update && \
    apt install -y openssh-server sudo && \
    mkdir /var/run/sshd

# create user
RUN useradd -m -s /bin/bash cbaciliod && \
    echo 'cbaciliod:cbaciliod' | chpasswd && \
    echo 'cbaciliod ALL=(ALL) ALL:ALL' >> /etc/sudoers
# Create the directory for the SSH daemon to run
# Copy your public key from the build context to the image
# Replace /path/to/your/public_key with the actual path in your build context
COPY keys/id_rsa.pub /tmp/id_rsa.pub
# Create SSH directory for the root user
RUN mkdir -p /home/cbaciliod/.ssh && \
    chmod 700 /home/cbaciliod/.ssh && \
    cat /tmp/id_rsa.pub >> /home/cbaciliod/.ssh/authorized_keys && \
    chmod 600 /home/cbaciliod/.ssh/authorized_keys && \
    chown -R cbaciliod:cbaciliod /home/cbaciliod/.ssh &&  \
    rm /tmp/id_rsa.pub

#RUN mkdir -p /root/.ssh && \
#    cat /tmp/id_rsa.pub >> /root/.ssh/authorized_keys && \
#    chmod 600 /root/.ssh/authorized_keys && \
#    rm /tmp/id_rsa.pub
#COPY keys/known_hosts /root/.ssh/known_hosts
#ADD sshd_config /etc/ssh/sshd_config

# Expose port 22 to the outside world
EXPOSE 22

# Start the SSH service
CMD ["/usr/sbin/sshd", "-D"]