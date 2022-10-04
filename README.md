# maximilian.lib.unb.ca
## Introduction
Maximilian is a ChatOps app allowing users to interact with Kubernetes resources via Slack.

## Slack Commands
* ```/drupal-uli``` : Provides a ULI link for Drupal instances

## k8s API Authentication
The application authenticates to the k8s API via two methods:
* It first attempts authentication using the configuration found at ```$HOME/.kube/config``` - this is useful for local development.
* If the above file does not exist, in-cluster API authentication (via a standard service account token mounted within deployed k8s pods) is attempted.
* If neither file exists, authentication will fail.

## Author / Contributors
This application was created at [![UNB Libraries](https://github.com/unb-libraries/assets/raw/master/unblibbadge.png "UNB Libraries")](https://lib.unb.ca) by the following humans:

<a href="https://github.com/JacobSanford"><img src="https://avatars.githubusercontent.com/u/244894?v=3" title="Jacob Sanford" width="128" height="128"></a>

## License
- As part of our 'open' ethos, UNB Libraries licenses its applications and workflows to be freely available to all whenever possible.
- Consequently, the contents of this repository [unb-libraries/maximilian.lib.unb.ca] are licensed under the [MIT License](http://opensource.org/licenses/mit-license.html). This license explicitly excludes:
  - Any website content, which remains the exclusive property of its author(s).
  - The UNB logo and any of the associated suite of visual identity assets, which remains the exclusive property of the University of New Brunswick.
