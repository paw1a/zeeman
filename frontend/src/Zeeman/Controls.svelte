<script lang="ts">
    import {field, form} from 'svelte-forms';
    import {min, required} from 'svelte-forms/validators';
    import {MagneticFieldDir} from "../api/zeeman";

    const resolution = field("resolution", "500", [required(), min(1)]);
    const pictureSize = field("pictureSize", "5", [required(), min(1)]);
    const waveLength = field("waveLength", "628.0", [required(), min(0)]);
    const focalDistance = field("focalDistance", "100", [required(), min(0)]);
    const glassesDistance = field("glassesDistance", "10", [required(), min(0)]);
    const pathDifference = field("pathDifference", "0", [required(), min(0)]);
    const refractionFactor = field("refractionFactor", "1.375", [required(), min(1)]);
    const magneticInduction = field("magneticInduction", "3", [required(), min(0)]);

    let magneticFieldDir = MagneticFieldDir.Parallel;

    const zeemanForm = form(resolution, pictureSize, waveLength, focalDistance, glassesDistance,
        pathDifference, refractionFactor, magneticInduction);

    const submitForm = e => {
        zeemanForm.validate();
    };

    const getDir = () => {
        console.log(magneticFieldDir);
        if (magneticFieldDir == MagneticFieldDir.Parallel) {
            return true;
        } else if (magneticFieldDir == MagneticFieldDir.Perpendicular) {
            return false;
        } else {
            console.error("Unexpected dir:", magneticFieldDir);
        }
    };

    const getMagneticPerpendicular = () => {
        console.log("perp", getDir());
        return !getDir() ? "active" : "";
    };

    const getMagneticParallel = () => {
        console.log("par", getDir());
        return getDir() ? "active" : "";
    };
</script>

<div class="col-md-4 order-md-2">
  <h4 style="margin-top: 2rem;">Параметры эксперимента</h4>

  <form class="needs-validation " novalidate="">
    <div class="mb-3">
      <label for="resolution">Разрешение изображения</label>
      <input id="resolution" class="form-control" type="text"
             bind:value={$resolution.value}>
    </div>
    <div class="mb-3">
      <label for="pictureSize">Размер пленки</label>
      <input id="pictureSize" class="form-control" type="text"
             bind:value={$pictureSize.value}>
    </div>
    <div class="mb-3">
      <label for="waveLength">Длина волны</label>
      <input id="waveLength" class="form-control" type="text"
             bind:value={$waveLength.value}>
    </div>
    <div class="mb-3">
      <label for="focalDistance">Фокусное расстояние</label>
      <input id="focalDistance" class="form-control" type="text"
             bind:value={$focalDistance.value}>
    </div>
    <div class="mb-3">
      <label for="glassesDistance">Расстояние между линзами</label>
      <input id="glassesDistance" class="form-control" type="text"
             bind:value={$glassesDistance.value}>
    </div>
    <div class="mb-3">
      <label for="pathDifference">Разность хода</label>
      <input id="pathDifference" class="form-control" type="text"
             bind:value={$pathDifference.value}>
    </div>
    <div class="mb-3">
      <label for="refractionFactor">Коэффициент преломления</label>
      <input id="refractionFactor" class="form-control" type="text"
             bind:value={$refractionFactor.value}>
    </div>
    <div class="mb-3">
      <label for="magneticInduction">Магнитная индукция</label>
      <input id="magneticInduction" class="form-control" type="text"
             bind:value={$magneticInduction.value}>
    </div>
    <div class="mb-3">
      <label>Направление магнитного поля</label>
      <div class="btn-group btn-group-toggle" data-toggle="buttons">
        <input class="btn-check magnetic-dir" type="radio" name="options" autocomplete="off"
               bind:group={magneticFieldDir} value={MagneticFieldDir.Parallel}
        id="dirParallel">
        <label class="btn btn-secondary" for="dirParallel">

          Параллельно</label>

        <input class="btn-check magnetic-dir" type="radio" name="options" autocomplete="off"
               bind:group={magneticFieldDir} value={MagneticFieldDir.Perpendicular}
              id="dirPerp">

        <label class="btn btn-secondary" for="dirPerp">
          Перпендикулярно
        </label>
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