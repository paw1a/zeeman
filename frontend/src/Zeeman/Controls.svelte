<div class="col-md-4 order-md-2">
  <h4 style="margin-top: 2rem;">Параметры эксперимента</h4>

  <form class="needs-validation " novalidate="">
    <div class="mb-3">
      <label for="waveLength">Длина волны</label>
      <input id="waveLength" class="form-control" type="text" placeholder="0.0" value="" required="">
    </div>
    <div class="mb-3">
      <label for="focalDistance">Фокусное расстояние</label>
      <input id="focalDistance" class="form-control" type="text" placeholder="0.0" value="" required="">
    </div>
    <div class="mb-3">
      <label for="glassesDistance">Расстояние до линз</label>
      <input id="glassesDistance" class="form-control" type="text" placeholder="0.0" value="" required="">
    </div>
    <div class="mb-3">
      <label for="pathDifference">Длина пути</label>
      <input id="pathDifference" class="form-control" type="text" placeholder="0.0" value="" required="">
    </div>
    <div class="mb-3">
      <label for="refractionFactor">Коэффициент преломления</label>
      <input id="refractionFactor" class="form-control" type="text" placeholder="0.0" value="" required="">
    </div>
    <div class="mb-3">
      <label for="magneticInduction">Магнитная индукция</label>
      <input id="magneticInduction" class="form-control" type="text" placeholder="0.0" value="" required="">
    </div>
    <button class="btn btn-primary" on:click={submitForm}>Показать результат</button>
  </form>
</div>

<script>
  const submitForm = e => {
  };
</script>