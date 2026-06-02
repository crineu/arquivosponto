#!/usr/bin/env sh

### parâmetros ############

BANCO="mysql"
BANCO_VERSAO="__8.0__"
VOLUME_LOCAL="${HOME}/pasta_desejada/${BANCO}_data"
NOME_MAQUINA="mysql-local"
PORTA="3306"
ROOT_PASSWORD="docker"
DATABASE="database"

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
  docker run \
    --rm \
    --name "$NOME_MAQUINA" \
    -p "$PORTA":3306 \
    -v "$VOLUME_LOCAL":/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD="$ROOT_PASSWORD" \
    -e MYSQL_DATABASE="$DATABASE" \
    "$BANCO:$BANCO_VERSAO" --default-authentication-plugin=mysql_native_password
}
    # --detach \


# Mysql >= 8 usa plugin de autenticação novo que pacote 'mysql' do node não suporta
# Por isso '--default-authentication-plugin=mysql_native_password' se mostra necessário
# Caso usuário tenha sido criado anteriormente, executar

# mysql -u root -p
# ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'docker';
# ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'docker';
# FLUSH PRIVILEGES;


info
inicializa_docker
