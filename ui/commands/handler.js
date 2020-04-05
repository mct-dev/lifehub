import { log } from '../logging/logger'

export function handleCommand (command) {
  const prefix = command.slice(0, 1)

  switch (prefix) {
    case Prefixes.Navigate:
      return handleNavigateCommand(command.slice(1))
    default:
      log(command)
  }
}

function handleNavigateCommand (command) {
  log(command)
}

const Prefixes = {
  Navigate: '/'
}
