<script lang="ts">
    import {MagneticFieldDir} from "../lib/api/zeeman";
    import {parameters} from "../lib/zeeman-form";
    import FormInput from "./FormInput.svelte";

    let magneticFieldDir = MagneticFieldDir.Parallel;

    let wasValidated = false;

    const submitForm = e => {
        wasValidated = true;
    };
</script>

<div class="col-md-4 order-md-2">
  <h4 style="margin-top: 2rem;">Параметры эксперимента</h4>

  <form class="needs-validation" class:was-validated={wasValidated} novalidate="">
    {#each parameters as param }
      <FormInput id="{param.id}" name="{param.name}" error="{param.error}" />
    {/each}
    <div class="mb-3">
      <label>Направление магнитного поля</label>
      <div class="btn-group btn-group-toggle" data-toggle="buttons">
        <input class="btn-check magnetic-dir" type="radio" name="options" autocomplete="off"
               bind:group={magneticFieldDir} value={MagneticFieldDir.Parallel}
               id="dirParallel">
        <label class="btn btn-secondary" for="dirParallel">Параллельно</label>

        <input class="btn-check magnetic-dir" type="radio" name="options" autocomplete="off"
               bind:group={magneticFieldDir} value={MagneticFieldDir.Perpendicular}
               id="dirPerp">
        <label class="btn btn-secondary" for="dirPerp">Перпендикулярно</label>
      </div>
    </div>
    <button class="btn btn-primary" on:click={submitForm}>Показать результат
    </button>
  </form>
</div>

<style>
    .magnetic-dir {
        appearance: none;
    }
</style>