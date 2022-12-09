<script lang="ts">
    import {MagneticFieldDir} from "../lib/api/zeeman";
    import {createFieldValidator, requiredValidator} from "../lib/validator";

    let resolution = "";
    let magneticFieldDir = MagneticFieldDir.Parallel;

    let wasValidated = false;

    const [validity, validate] = createFieldValidator(requiredValidator());

    const submitForm = e => {
        wasValidated = true;
    };
</script>

<div class="col-md-4 order-md-2">
  <h4 style="margin-top: 2rem;">Параметры эксперимента</h4>

  <form class="needs-validation" class:was-validated={wasValidated} novalidate="">
    <div class="mb-3">
      <label class="form-label" for="resolution">Разрешение изображения</label>
      <input id="resolution"
             class="form-control"
             type="text"
             class:is-invalid={!$validity.valid}
             class:is-valid={$validity.valid}
             bind:value={resolution}
             use:validate={resolution}
      >
      {#if !$validity.valid}
      <div class="invalid-feedback">
         Введите правильное разрешение
      </div>
      {/if}
    </div>
<!--    <div class="mb-3">-->
<!--      <label for="pictureSize">Размер пленки</label>-->
<!--      <input id="pictureSize" class="form-control" type="text"-->
<!--             bind:value={$pictureSize.value}>-->
<!--    </div>-->
<!--    <div class="mb-3">-->
<!--      <label for="waveLength">Длина волны</label>-->
<!--      <input id="waveLength" class="form-control" type="text"-->
<!--             bind:value={$waveLength.value}>-->
<!--    </div>-->
<!--    <div class="mb-3">-->
<!--      <label for="focalDistance">Фокусное расстояние</label>-->
<!--      <input id="focalDistance" class="form-control" type="text"-->
<!--             bind:value={$focalDistance.value}>-->
<!--    </div>-->
<!--    <div class="mb-3">-->
<!--      <label for="glassesDistance">Расстояние между линзами</label>-->
<!--      <input id="glassesDistance" class="form-control" type="text"-->
<!--             bind:value={$glassesDistance.value}>-->
<!--    </div>-->
<!--    <div class="mb-3">-->
<!--      <label for="pathDifference">Разность хода</label>-->
<!--      <input id="pathDifference" class="form-control" type="text"-->
<!--             bind:value={$pathDifference.value}>-->
<!--    </div>-->
<!--    <div class="mb-3">-->
<!--      <label for="refractionFactor">Коэффициент преломления</label>-->
<!--      <input id="refractionFactor" class="form-control" type="text"-->
<!--             bind:value={$refractionFactor.value}>-->
<!--    </div>-->
<!--    <div class="mb-3">-->
<!--      <label for="magneticInduction">Магнитная индукция</label>-->
<!--      <input id="magneticInduction" class="form-control" type="text"-->
<!--             bind:value={$magneticInduction.value}>-->
<!--    </div>-->
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