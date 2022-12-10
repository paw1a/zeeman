<svelte:options accessors={true}/>

<script>
    import {createFieldValidator, requiredValidator} from "../lib/validator.ts";

    export let id;
    export let name;
    export let error;
    export let init;

    const [validity, validate] = createFieldValidator(requiredValidator());

    export let value = init;
</script>

<div class="mb-3">
  <label class="form-label" for="{id}">{name}</label>
  <input id="{id}"
         class="form-control"
         type="text"
         class:is-invalid={!$validity.valid}
         class:is-valid={$validity.valid}
         bind:value={value}
         use:validate={value}
  />

  {#if !$validity.valid}
      <div class="invalid-feedback">
        {error}
      </div>
    {/if}

</div>