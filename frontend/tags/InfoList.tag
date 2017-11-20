<InfoList>
  <virtual each="{ list }">
    <Info opts="{ this }"></Info>
  </virtual>

  <script>
   request = window.superagent
   var self = this

   request.get("/api")
          .end(function(err, res) {
            self.list = Array.prototype.reverse(res.body.info)
            self.update()
          })
  </script>
</InfoList>
