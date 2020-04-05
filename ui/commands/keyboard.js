import { log } from '../logging/logger'

/**
 * @param {string[]} combo
 */
export function handleKeyCombo (combo) {
  const foundCombo = allowedKeyCombinations.find((ac) => {
    let isMatchingCombo = false

    if (ac.length === combo.length) {
      ac.forEach((key, ix) => {
        isMatchingCombo = key === combo[ix]
      })
    }

    return isMatchingCombo
  })

  // TODO: do something other than log the key combo
  log('foundCombo: ', foundCombo)
}

const allowedKeyCombinations = [
  ['Shift', ' ']
]
