#!/bin/bash
if ! command -v envy &> /dev/null
then
  if ! command -v go &> /dev/null
  then
    echo "No installations of envy or Go found on system. Skipping hook gracefully"
    exit 0
  fi
  go install github.com/softmaxer/envy
fi


echo $PATH | grep -q "$HOME/go/bin"
if [$? -eq 1]; then
  export PATH="$PATH:$HOME/go/bin/"
fi

envy new
envy pack
