<script lang="ts">
    let className = "";
    let iconName = "";
    let iconType: "filled" | "outlined" | "round" | "sharp" | "two-tone" = "outlined";
    export { className as class, iconName as name, iconType as type };
    let material = `material-icons-${iconType}`;
  </script>
  <i translate="no" class="select-none {material} {className}">{iconName}</i>
  
  <style global>
    @font-face {
      font-family: "Material Icons";
      font-style: normal;
      font-weight: 400;
      src: url(https://fonts.gstatic.com/s/materialicons/v139/flUhRq6tzZclQEJ-Vdg-IuiaDsNc.woff2) format("woff2");
    }
    .select-none {
      user-select: none;
    }
    [class^="material-icons-"],
    [class*=" material-icons-"] {
      font-family: "Material Icons";
      font-weight: normal;
      font-style: normal;
      font-size: 24px;
      line-height: 1;
      letter-spacing: normal;
      text-transform: none;
      display: inline-block;
      white-space: nowrap;
      word-wrap: normal;
      direction: ltr;
      font-feature-settings: "liga";
      -webkit-font-smoothing: antialiased;
    }
  </style>