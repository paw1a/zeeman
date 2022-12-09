<script lang="ts">
    import zeemanAPIRequest, {MagneticFieldDir} from "../lib/api/zeeman";
    import {parameters} from "../lib/zeeman-form";
    import FormInput from "./FormInput.svelte";

    let magneticFieldDir = MagneticFieldDir.Parallel;

    let params = {};
    const submitForm = e => {
        let data = {};
        for (let param in params) {
            data[param] = params[param].value;
        }
        data["magneticFieldDirection"] = magneticFieldDir == MagneticFieldDir.Parallel ? 1 : 2;
        let dto = {
            resolution: parseFloat(data["resolution"]),
            pictureSize: parseFloat(data["pictureSize"]),
            waveLength: parseFloat(data["waveLength"]),
            focalDistance: parseFloat(data["focalDistance"]),
            glassesDistance: parseFloat(data["glassesDistance"]),
            pathDifference: parseFloat(data["pathDifference"]),
            refractionFactor: parseFloat(data["refractionFactor"]),
            magneticInduction: parseFloat(data["magneticInduction"]),
            magneticFieldDirection: parseFloat(data["magneticFieldDirection"])
        };
        zeemanAPIRequest(dto);
    };
</script>

<div class="col-md-4 order-md-2">
  <h4 style="margin-top: 2rem;">Параметры эксперимента</h4>

  <form class="needs-validation"
        novalidate=""
        on:submit|preventDefault={submitForm}
  >
    {#each parameters as param }
      <FormInput id="{param.id}" name="{param.name}" error="{param.error}" init="{param.init}"
                 bind:this={params[param.id]}/>
    {/each}
    <div class="mb-3">
      <label>Направление магнитного поля</label>
      <div class="btn-group btn-group-toggle" data-toggle="buttons">
        <input class="btn-check" type="radio" name="options" autocomplete="off"
               bind:group={magneticFieldDir} value={MagneticFieldDir.Parallel}
               id="dirParallel">
        <label class="btn btn-secondary" for="dirParallel">Параллельно</label>

        <input class="btn-check" type="radio" name="options" autocomplete="off"
               bind:group={magneticFieldDir} value={MagneticFieldDir.Perpendicular}
               id="dirPerp">
        <label class="btn btn-secondary" for="dirPerp">Перпендикулярно</label>
      </div>
    </div>
    <input type="submit" class="btn btn-primary" value="Показать результат"/>
  </form>
</div>