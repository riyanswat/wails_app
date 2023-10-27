<script>
  import logo from "./assets/images/logo-universal.png";
  import { GeneratePassword } from "../wailsjs/go/main/App.js";

  let resultText = "Generated Password: ";
  let length = 12; // Default password length
  let includeLowercase = true;
  let includeUppercase = true;
  let includeNumbers = true;
  let includeSpecialChars = true;

  function generatePassword() {
    GeneratePassword(
      length,
      includeLowercase,
      includeUppercase,
      includeNumbers,
      includeSpecialChars
    ).then((result) => (resultText = "Generated Password: " + result));
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src={logo} />
  <div class="result" id="result">{resultText}</div>
  <div class="password-options">
    <label for="length">Password Length:</label>
    <input type="number" bind:value={length} min="4" max="64" />
    <label for="includeLowercase">Include Lowercase:</label>
    <input type="checkbox" bind:checked={includeLowercase} />
    <label for="includeUppercase">Include Uppercase:</label>
    <input type="checkbox" bind:checked={includeUppercase} />
    <label for="includeNumbers">Include Numbers:</label>
    <input type="checkbox" bind:checked={includeNumbers} />
    <label for="includeSpecialChars">Include Special Characters:</label>
    <input type="checkbox" bind:checked={includeSpecialChars} />
  </div>
  <button class="btn" on:click={generatePassword}>generate password</button>
</main>

<!-- ========================================================= -->

<!-- <script>
  import logo from "./assets/images/logo-universal.png";
  import { Greet } from "../wailsjs/go/main/App.js";

  let resultText = "Please enter your name below 👇";
  let name;

  function greet() {
    Greet(name).then((result) => (resultText = result));
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src={logo} />
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input
      autocomplete="off"
      bind:value={name}
      class="input"
      id="name"
      type="text"
    />
    <button class="btn" on:click={greet}>greet</button>
  </div>
</main>

<style>
  #logo {
    display: block;
    width: 30%;
    height: 25%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
    border: none;
    border-radius: 50%;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
    font-variant: small-caps;
    transition: background-color 0.3s ease, color 0.3s ease;
  }

  .input-box .btn:hover {
    /* background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%); */
    color: #ffffff;
    background-color: #000811;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    /* padding: 0 10px; */
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style> -->

<style>
  /* ... (your existing CSS) ... */

  .password-options {
    display: flex;
    flex-wrap: wrap;
    margin: 1rem auto;
    justify-content: space-around;
  }

  .password-options label {
    font-weight: bold;
  }

  .password-options input[type="number"] {
    width: 50px;
  }
</style>
