PORT="${BPL_DEBUG_PORT:=8000}"
SUSPEND="${BPL_DEBUG_SUSPEND:=n}"

printf "Debugging enabled on port %s" "${PORT}"
if [[ "${SUSPEND}" = "y" ]]; then
  printf ", suspended on start"
fi
printf "\n"

export JAVA_OPTS="${JAVA_OPTS} -agentlib:jdwp=transport=dt_socket,server=y,address=${PORT},suspend=${SUSPEND}"
