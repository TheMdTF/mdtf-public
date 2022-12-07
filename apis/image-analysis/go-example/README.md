# Sample Image Analysis System

## Running and Testing via Docker
*PREREQUISITES*:
 * A clone of the git repository
 * An installation of [docker](https://docs.docker.com/install/#supported-platforms)

To run the sample application via docker:
 * cd `./docker`
 * package the image using `./make-docker-image.sh`
 * run the image using `./run-docker-image.sh`

After the steps above, the server should be accessible at `localhost:8080`.  The script `./docker/save-docker-image.sh` will output the docker image in the format we expect for delivery to the MdTF.  Please utilize the scripts' `[COMPANY_NAME]` command line parameter when preparing for delivery.
