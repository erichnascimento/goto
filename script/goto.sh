#!/usr/bin/env bash

#
# include this file in your ~/.bashrc
#

#
# Try use goto program from env var
# otherwise search in path
# 
GOTO_BIN=${GOTO_BIN:-`which gotos`}

#
# Wrapper function for cd command
# 
goto () {
  if [ ! -x "$GOTO_BIN" ]; then
    echo "goto program not found"
    return 1
  fi

  # When user try go to a path (goto <name>)...
  if [ $# = 1 ] && [ $@ != "list" ] && [[ "$@" =~ ^[^-] ]] ; then
    GOTO_=`$GOTO_BIN $@`
    if [ ! -z $GOTO_ ]; then
      cd $GOTO_
      return 0
    else 
      echo "  Sorry, path entry \"$1\" not found..."
      return 1
    fi
  fi

  # Clear result var
  GOTO_=

  $GOTO_BIN "$@"
  return $?

}