#!/usr/bin/env sh

### parâmetros ############

# Para funcionar o firewall do ubuntu (ufw) é ncessário utilizar
# o modo `--network host` do docker
# Portanto a variável PORTA é ignorada e será sempre 5432

BANCO="postgres"
BANCO_VERSAO="11-alpine"
VOLUME_LOCAL="${HOME}/postgres/${BANCO}_data"
NOME_MAQUINA="pg-90.20"
PORTA="5432"
ROOT_USER="docker"
ROOT_PASSWORD="docker"
DATABASE="sistemasara"

###########################


set -o errexit  # exit on error
set -o nounset  # don't allow unset variables
# set -o xtrace # enable for debugging

reset=$(tput sgr0)
gray=$(tput setaf  248)
red=$(tput setaf   160)
green=$(tput setaf 76)
blue=$(tput setaf  45)

e_sucesso() { printf "${green}[ok] %s${reset}\n" "$@"; }
e_erro()    { printf "${red}[!erro!] %s${reset}\n" "$@"; }
e_aviso()   { printf "${gray}%s${reset}\n" "$@"; }

# variáveis úteis
# __dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# __file="${__dir}/$(basename "${BASH_SOURCE[0]}")"
# __base="$(basename ${__file} .sh)"
# __root="$(cd "$(dirname "${__dir}")" && pwd)"


info() {
  e_aviso "Iniciando máquina DOCKER em background [$(basename $0)]"
  printf  "       Imagem:  "; e_sucesso "$BANCO:$BANCO_VERSAO (porta $PORTA)"
  printf  " Volume_local:  "; verifica_volume_local
  printf  " Desligamento:  "; e_aviso "docker stop $NOME_MAQUINA"
  printf  "\n"
}

verifica_volume_local() {
  if [ -d "$VOLUME_LOCAL" ]
  then
      e_sucesso "$VOLUME_LOCAL"
  else
      e_erro "$VOLUME_LOCAL"
      printf "\n"
      printf "mkdir $VOLUME_LOCAL"
      printf "\n\n"
      exit 1
  fi
}

inicializa_docker() {
  docker run -d \
    --name "$NOME_MAQUINA" \
    --network host \
    --restart=unless-stopped \
    -v "$VOLUME_LOCAL":/var/lib/postgresql/data \
    -e POSTGRES_USER="$ROOT_USER" \
    -e POSTGRES_PASSWORD="$ROOT_PASSWORD" \
    -e POSTGRES_DB="$DATABASE" \
    "$BANCO:$BANCO_VERSAO" 
}
#    -p "$PORTA":5432 \

info
inicializa_docker
