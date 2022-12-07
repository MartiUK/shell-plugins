<p align="center">
  <a href="https://1password.com">
    <img alt="1Password logo" src="https://user-images.githubusercontent.com/7430639/205342015-46801fd8-6701-482f-9da9-e21e7e39b3a1.svg" height="64">
    <h3 align="center">1Password Shell Plugins</h3>
  </a>
</p>

<p align="center">
  Seamless authentication for every tool in your terminal. <sup><b><a href="#-beta-notice">[BETA]</a></b></sup>
</p>

<p align="center">
  <a href="https://developer.1password.com/docs/cli/shell-plugins"><b>Available plugins</b></a> ·
  <a href="https://developer.1password.com/docs/cli/"><b>1Password CLI docs</b></a> ·
  <a href="https://developer.1password.com/docs/cli/shell-plugins/contribute"><b>Contribution docs</b></a>
</p>
<br/>

---

No more plaintext credentials in your home directory. Automatically authenticate every CLI you use with [1Password CLI](https://developer.1password.com/docs/cli/) + [Shell Plugins](https://developer.1password.com/docs/cli/shell-plugins/). Approve temporary credential usage in your terminal with biometrics.

## 🪄 Usage

 <picture>
   <source srcset="https://user-images.githubusercontent.com/1965218/206192140-cfa3103e-701c-46c1-9d7c-bd6a404b8823.gif">
   <img src="https://user-images.githubusercontent.com/7430639/205337855-41604aca-0ddb-4eab-a5f0-fb9107e09d8d.gif" alt="Example of 1Password Shell Plugins with AWS: user runs an `aws` command, a Touch ID prompt shows up, and `aws` is automatically authenticated" style="max-width: 100%; display: inline-block;">
</picture>

## 🚀 Get started

* Get started with [1Password Shell Plugins](https://developer.1password.com/docs/cli/shell-plugins)
* Get started with [creating your own plugins](https://developer.1password.com/docs/cli/shell-plugins/contribute) <sup><b>[BETA]</b></sup>

## 👫 Contributing

This repository contains the list of available plugins, as well as the SDK to create new plugins.

Is your favorite CLI not listed yet? Learn how to [build a shell plugin](https://developer.1password.com/docs/cli/shell-plugins/contribute) yourself and [open a PR](https://github.com/1Password/shell-plugins/pulls)!

If you need help when building a plugin, create an issue on GitHub or join our [Developer Slack workspace](https://join.slack.com/t/1password-devs/shared_invite/zt-1halo11ps-6o9pEv96xZ3LtX_VE0fJQA) to tell us about your plugin proposal and we can advise you on the most suitable approach for your use case.

For the contribution guidelines, see [CONTRIBUTING.md](CONTRIBUTING.md).

## 📣 Beta Notice

The plugin ecosystem is still in beta. In practice, this means that if you're locally building your own plugins, you'll likely have to recompile your plugin every now and then to keep up with the latest updates of the 1Password CLI.
