<!doctype html>
<html lang="ja">
  <head>
    <meta charset="UTF-8"/>
    <title>Information::charakoba.com</title>
    {% include "include.tpl" %}
    <script>
     $(document).ready(function() {
       $('select').material_select();
       $('.datepicker').pickadate({
         selectMonths: true,
         selectYears: 5,
         format: 'yyyy/mm/dd'
       });
       $('textarea#detail').characterCounter();
       $('#submit').click(post_info);
     });
     function post_info() {
       $.post(
         "/api",
         {
           "type": $('#type').val(),
           "service": $('#service').val(),
           "begin": $('#begin-date').val()+' '+$('#begin-hour').val()+':'+$('#begin-minute').val()+':'+$('#begin-second').val(),
           "end": $('#end-date').val()+' '+$('#end-hour').val()+':'+$('#end-minute').val()+':'+$('#end-second').val(),
           "detail": $('#detail').val(),
           "apikey": $('#apikey').val()
         },
         function(){
           $(location).attr('href', '/');
         }
       );
     }
    </script>
  </head>
  <body>
    {% include "navbar.tpl" %}
    <main class="container">
      <div class="row">
        <div class="col s10 offset-s1">
          <form class="card grey lighten-5">
            <div class="card-content">
              <span class="card-title">POST INFORMATION</span>
              <div class="container">
                <div class="row">
                  <div class="input-field col s12">
                    <label for="">type</label>
                  </div>
                </div>
                <div class="row">
                  <div class="input-field col s12">
                    <select id="type" name="type">
                      <option value="event">event</option>
                      <option value="maintenance">maintenance</option>
                    </select>
                  </div>
                </div>
                <div class="row">
                  <div class="input-field col s12">
                    <label for="">service</label>
                  </div>
                </div>
                <div class="row">
                  <div class="input-field col s12">
                    <select id="service" name="service">
                      {% for service in services %}
                      <option value="{{ service }}">{{ service }}</option>
                      {% endfor %}
                    </select>
                  </div>
                </div>
                <div class="row">
                  <div class="input-field col s12">
                    <label for="">begin</label>
                  </div>
                </div>
                <div class="row">
                  <div class="ipnut-field col s12">
                    <input class="datepicker" id="begin-date" name="begin-date" type="date" value="" type-date=""/>
                  </div>
                </div>
                <div class="row">
                  <div class="input-field col s2">
                    <select id="begin-hour" name="begin-hour">
                      <option value="00">0時</option>
                      <option value="01">1時</option>
                      <option value="02">2時</option>
                      <option value="03">3時</option>
                      <option value="04">4時</option>
                      <option value="05">5時</option>
                      <option value="06">6時</option>
                      <option value="07">7時</option>
                      <option value="08">8時</option>
                      <option value="09">9時</option>
                      <option value="10">10時</option>
                      <option value="11">11時</option>
                      <option value="12">12時</option>
                      <option value="13">13時</option>
                      <option value="14">14時</option>
                      <option value="15">15時</option>
                      <option value="16">16時</option>
                      <option value="17">17時</option>
                      <option value="18">18時</option>
                      <option value="19">19時</option>
                      <option value="20">20時</option>
                      <option value="21">21時</option>
                      <option value="22">22時</option>
                      <option value="23">23時</option>
                    </select>
                  </div>
                  <div class="input-field col s2">
                    <select id="begin-minute" name="begin-minute">
                      <option value="00">0分</option>
                      <option value="01">1分</option>
                      <option value="02">2分</option>
                      <option value="03">3分</option>
                      <option value="04">4分</option>
                      <option value="05">5分</option>
                      <option value="06">6分</option>
                      <option value="07">7分</option>
                      <option value="08">8分</option>
                      <option value="09">9分</option>
                      <option value="10">10分</option>
                      <option value="11">11分</option>
                      <option value="12">12分</option>
                      <option value="13">13分</option>
                      <option value="14">14分</option>
                      <option value="15">15分</option>
                      <option value="16">16分</option>
                      <option value="17">17分</option>
                      <option value="18">18分</option>
                      <option value="19">19分</option>
                      <option value="20">20分</option>
                      <option value="21">21分</option>
                      <option value="22">22分</option>
                      <option value="23">23分</option>
                      <option value="24">24分</option>
                      <option value="25">25分</option>
                      <option value="26">26分</option>
                      <option value="27">27分</option>
                      <option value="28">28分</option>
                      <option value="29">29分</option>
                      <option value="30">30分</option>
                      <option value="31">31分</option>
                      <option value="32">32分</option>
                      <option value="33">33分</option>
                      <option value="34">34分</option>
                      <option value="35">35分</option>
                      <option value="36">36分</option>
                      <option value="37">37分</option>
                      <option value="38">38分</option>
                      <option value="39">39分</option>
                      <option value="40">40分</option>
                      <option value="41">41分</option>
                      <option value="42">42分</option>
                      <option value="43">43分</option>
                      <option value="44">44分</option>
                      <option value="45">45分</option>
                      <option value="46">46分</option>
                      <option value="47">47分</option>
                      <option value="48">48分</option>
                      <option value="49">49分</option>
                      <option value="50">50分</option>
                      <option value="51">51分</option>
                      <option value="52">52分</option>
                      <option value="53">53分</option>
                      <option value="54">54分</option>
                      <option value="55">55分</option>
                      <option value="56">56分</option>
                      <option value="57">57分</option>
                      <option value="58">58分</option>
                      <option value="59">59分</option>
                    </select>
                  </div>
                  <div class="input-field col s2">
                    <select id="begin-second" name="begin-second">
                      <option value="00">0秒</option>
                      <option value="01">1秒</option>
                      <option value="02">2秒</option>
                      <option value="03">3秒</option>
                      <option value="04">4秒</option>
                      <option value="05">5秒</option>
                      <option value="06">6秒</option>
                      <option value="07">7秒</option>
                      <option value="08">8秒</option>
                      <option value="09">9秒</option>
                      <option value="10">10秒</option>
                      <option value="11">11秒</option>
                      <option value="12">12秒</option>
                      <option value="13">13秒</option>
                      <option value="14">14秒</option>
                      <option value="15">15秒</option>
                      <option value="16">16秒</option>
                      <option value="17">17秒</option>
                      <option value="18">18秒</option>
                      <option value="19">19秒</option>
                      <option value="20">20秒</option>
                      <option value="21">21秒</option>
                      <option value="22">22秒</option>
                      <option value="23">23秒</option>
                      <option value="24">24秒</option>
                      <option value="25">25秒</option>
                      <option value="26">26秒</option>
                      <option value="27">27秒</option>
                      <option value="28">28秒</option>
                      <option value="29">29秒</option>
                      <option value="30">30秒</option>
                      <option value="31">31秒</option>
                      <option value="32">32秒</option>
                      <option value="33">33秒</option>
                      <option value="34">34秒</option>
                      <option value="35">35秒</option>
                      <option value="36">36秒</option>
                      <option value="37">37秒</option>
                      <option value="38">38秒</option>
                      <option value="39">39秒</option>
                      <option value="40">40秒</option>
                      <option value="41">41秒</option>
                      <option value="42">42秒</option>
                      <option value="43">43秒</option>
                      <option value="44">44秒</option>
                      <option value="45">45秒</option>
                      <option value="46">46秒</option>
                      <option value="47">47秒</option>
                      <option value="48">48秒</option>
                      <option value="49">49秒</option>
                      <option value="50">50秒</option>
                      <option value="51">51秒</option>
                      <option value="52">52秒</option>
                      <option value="53">53秒</option>
                      <option value="54">54秒</option>
                      <option value="55">55秒</option>
                      <option value="56">56秒</option>
                      <option value="57">57秒</option>
                      <option value="58">58秒</option>
                      <option value="59">59秒</option>
                    </select>
                  </div>
                </div>
                <div class="row">
                  <div class="input-field col s12">
                    <label for="">end</label>
                  </div>
                </div>
                <div class="row">
                  <div class="ipnut-field col s12">
                    <input class="datepicker" id="end-date" name="end-date" type="date" value="" type-date=""/>
                  </div>
                </div>
                <div class="row">
                  <div class="input-field col s2">
                    <select id="end-hour" name="end-hour">
                      <option value="00">0時</option>
                      <option value="01">1時</option>
                      <option value="02">2時</option>
                      <option value="03">3時</option>
                      <option value="04">4時</option>
                      <option value="05">5時</option>
                      <option value="06">6時</option>
                      <option value="07">7時</option>
                      <option value="08">8時</option>
                      <option value="09">9時</option>
                      <option value="10">10時</option>
                      <option value="11">11時</option>
                      <option value="12">12時</option>
                      <option value="13">13時</option>
                      <option value="14">14時</option>
                      <option value="15">15時</option>
                      <option value="16">16時</option>
                      <option value="17">17時</option>
                      <option value="18">18時</option>
                      <option value="19">19時</option>
                      <option value="20">20時</option>
                      <option value="21">21時</option>
                      <option value="22">22時</option>
                      <option value="23">23時</option>
                    </select>
                  </div>
                  <div class="input-field col s2">
                    <select id="end-minute" name="end-minute">
                      <option value="00">0分</option>
                      <option value="01">1分</option>
                      <option value="02">2分</option>
                      <option value="03">3分</option>
                      <option value="04">4分</option>
                      <option value="05">5分</option>
                      <option value="06">6分</option>
                      <option value="07">7分</option>
                      <option value="08">8分</option>
                      <option value="09">9分</option>
                      <option value="10">10分</option>
                      <option value="11">11分</option>
                      <option value="12">12分</option>
                      <option value="13">13分</option>
                      <option value="14">14分</option>
                      <option value="15">15分</option>
                      <option value="16">16分</option>
                      <option value="17">17分</option>
                      <option value="18">18分</option>
                      <option value="19">19分</option>
                      <option value="20">20分</option>
                      <option value="21">21分</option>
                      <option value="22">22分</option>
                      <option value="23">23分</option>
                      <option value="24">24分</option>
                      <option value="25">25分</option>
                      <option value="26">26分</option>
                      <option value="27">27分</option>
                      <option value="28">28分</option>
                      <option value="29">29分</option>
                      <option value="30">30分</option>
                      <option value="31">31分</option>
                      <option value="32">32分</option>
                      <option value="33">33分</option>
                      <option value="34">34分</option>
                      <option value="35">35分</option>
                      <option value="36">36分</option>
                      <option value="37">37分</option>
                      <option value="38">38分</option>
                      <option value="39">39分</option>
                      <option value="40">40分</option>
                      <option value="41">41分</option>
                      <option value="42">42分</option>
                      <option value="43">43分</option>
                      <option value="44">44分</option>
                      <option value="45">45分</option>
                      <option value="46">46分</option>
                      <option value="47">47分</option>
                      <option value="48">48分</option>
                      <option value="49">49分</option>
                      <option value="50">50分</option>
                      <option value="51">51分</option>
                      <option value="52">52分</option>
                      <option value="53">53分</option>
                      <option value="54">54分</option>
                      <option value="55">55分</option>
                      <option value="56">56分</option>
                      <option value="57">57分</option>
                      <option value="58">58分</option>
                      <option value="59">59分</option>
                    </select>
                  </div>
                  <div class="input-field col s2">
                    <select id="end-second" name="end-second">
                      <option value="00">0秒</option>
                      <option value="01">1秒</option>
                      <option value="02">2秒</option>
                      <option value="03">3秒</option>
                      <option value="04">4秒</option>
                      <option value="05">5秒</option>
                      <option value="06">6秒</option>
                      <option value="07">7秒</option>
                      <option value="08">8秒</option>
                      <option value="09">9秒</option>
                      <option value="10">10秒</option>
                      <option value="11">11秒</option>
                      <option value="12">12秒</option>
                      <option value="13">13秒</option>
                      <option value="14">14秒</option>
                      <option value="15">15秒</option>
                      <option value="16">16秒</option>
                      <option value="17">17秒</option>
                      <option value="18">18秒</option>
                      <option value="19">19秒</option>
                      <option value="20">20秒</option>
                      <option value="21">21秒</option>
                      <option value="22">22秒</option>
                      <option value="23">23秒</option>
                      <option value="24">24秒</option>
                      <option value="25">25秒</option>
                      <option value="26">26秒</option>
                      <option value="27">27秒</option>
                      <option value="28">28秒</option>
                      <option value="29">29秒</option>
                      <option value="30">30秒</option>
                      <option value="31">31秒</option>
                      <option value="32">32秒</option>
                      <option value="33">33秒</option>
                      <option value="34">34秒</option>
                      <option value="35">35秒</option>
                      <option value="36">36秒</option>
                      <option value="37">37秒</option>
                      <option value="38">38秒</option>
                      <option value="39">39秒</option>
                      <option value="40">40秒</option>
                      <option value="41">41秒</option>
                      <option value="42">42秒</option>
                      <option value="43">43秒</option>
                      <option value="44">44秒</option>
                      <option value="45">45秒</option>
                      <option value="46">46秒</option>
                      <option value="47">47秒</option>
                      <option value="48">48秒</option>
                      <option value="49">49秒</option>
                      <option value="50">50秒</option>
                      <option value="51">51秒</option>
                      <option value="52">52秒</option>
                      <option value="53">53秒</option>
                      <option value="54">54秒</option>
                      <option value="55">55秒</option>
                      <option value="56">56秒</option>
                      <option value="57">57秒</option>
                      <option value="58">58秒</option>
                      <option value="59">59秒</option>
                    </select>
                  </div>
                </div>
                <div class="row">
                  <div class="col s12 input-field">
                    <label for="">detail</label>
                  </div>
                </div>
                <div class="row">
                  <div class="input-field col s12">
                    <textarea id="detail" class="materialize-textarea" cols="30" name="detail" rows="10" length="128"></textarea>
                  </div>
                </div>
                <div class="row">
                  <div class="input-field col s12">
                    <label for="apikey">api-key</label>

                    <input id="apikey" name="apikey" type="text" value=""/>
                  </div>
                </div>
              </div>
            </div>
            <div class="card-action">
              <a href="/" class="btn">Cancel</a>
              <a class="btn" id="submit" name="submit">POST <i class="material-icons right">send</i></a>
            </div>
          </form>
        </div>
      </div>
    </main>
  </body>
</html>
