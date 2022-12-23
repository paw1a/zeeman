import {writable} from 'svelte/store'

export function buildValidator(validators) {
    return function validate(value, dirty) {
        if (!validators || validators.length === 0) {
            return {dirty, valid: true}
        }

        const failing = validators.find(v => v(value) !== true)

        return {
            dirty,
            valid: !failing,
            message: failing && failing(value)
        }
    }
}

export function createFieldValidator(...validators) {
    const {subscribe, set} = writable({dirty: false, valid: false, message: null})
    const validator = buildValidator(validators)

    function action(node, binding) {
        function validate(value, dirty) {
            const result = validator(value, dirty)
            set(result)
        }

        validate(binding, false)

        return {
            update(value) {
                validate(value, true)
            }
        }
    }

    return [{subscribe}, action]
}

function requiredValidator() {
    return function required(value) {
        return (value !== undefined && value !== null && value !== '') || 'This field is required'
    }
}

export {
    emailValidator,
    requiredValidator
}
