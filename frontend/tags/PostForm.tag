


<PostForm>
  <form onsubmit="{ submitInfo }">
    <fieldset class="uk-fieldset">
      <div class="uk-margin">
        <select class="uk-select">
          <option each="{ types }">{ type }</option>
        </select>
      </div>
      <div class="uk-margin">
        <select class="uk-select">
          <option each="{ services}">{ name }</option>
        </select>
      </div>
      <div class="uk-margin">
        <label for="begin">begin:</label>
        <input class="uk-input" name="begin" id="begin" type="date" placeholder="YYYY-MM-DD hh:mm:ss"/>
      </div>
      <div class="uk-margin">
        <label for="end">end:</label>
        <input class="uk-input" name="end" id="end" type="date" placeholder="YYYY-MM-DD hh:mm:ss"/>
      </div>
      <div class="uk-margin">
        <textarea class="uk-textarea" rows="5" placeholder="Detail" name="detail"></textarea>
      </div>
      <div class="uk-margin">
        <input class="uk-input" type="text" placeholder="Token" name="token"/>
      </div>
      <p class="uk-text-right">
        <button class="uk-button uk-button-default uk-modal-close" type="button">Cancel</button>
        <button class="uk-button uk-button-primary" type="submit">post</button>
      </p>
    </fieldset>
  </form>

  <script>
   var request = window.superagent
   var self = this
   request.get("/types")
          .end(function(err, res) {
            self.types = res.body.types
            self.update()
          })
   request.get("/services")
          .end(function(err, res) {
            self.services = res.body.services
            self.update()
          })

   submitInfo(e) {
     var newInfo = { info_type: e.target.1.value, service: e.target.2.value, begin: e.target.3.value, token: e.target.6.value }
     if (e.target.4.value != "") {
       newInfo.end = e.target.4.value
     }
     if (e.target.5.value != "") {
       newInfo.detail = e.target.5.value
     }

     request.post("/")
            .send(newInfo)
            .end(function(err, res) {})
   }
  </script>
</PostForm>
