FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    curl \
    wget \
    jq \
    vim

RUN curl -LO "http://repo.continuum.io/miniconda/Miniconda3-latest-Linux-x86_64.sh"
RUN bash Miniconda3-latest-Linux-x86_64.sh -p /miniconda -b
RUN rm Miniconda3-latest-Linux-x86_64.sh
ENV PATH=/miniconda/bin:${PATH}
RUN pip install jupyter
RUN pip install pandas
RUN conda update --all --yes
RUN printf "UID=$(id -u)\nGID=$(id -g)\n" > .env
CMD ["jupyter", "notebook", "--port=8888", "--no-browser", "--ip=0.0.0.0", "--allow-root"]
