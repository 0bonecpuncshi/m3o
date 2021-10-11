<img src="images/banner.png" />

# M3O

[M3O](https://m3o.com) is an open source public cloud platform. We are building an AWS alternative for the next generation of developers.

## Overview

AWS has becoming the monstrous behemoth public cloud provider we all love to hate. It's massive complexity and confusing pricing has made it an incredibly 
difficult choice for developers to adopt. But what's the alternative? M3O is a new open source public cloud platform for the next generation of developers. Consume existing public APIs all in one place as simpler programmable building blocks. 

M3O is powered by the open source [Micro](https://github.com/micro/micro) platform and programmable real world [Micro Services](https://github.com/micro/services).

## Features

- **🔥 Dev UX** - The developer experience is first priority. A slick new UX for the next generation of developers.
- **☝️ One Token** - Use one Micro API token to fulfill all your API needs. Access multiple public APIs with a single token.
- **⚡ Fast Access** - Using a new API is easy - no need to learn yet another API, it's all the same Micro developer experience.
- **🆓 Free to start** - It's a simple pay as you go model and everything is priced per request. Top up your account and start making calls.
- **🚫 Anti AWS Billing** - Don't get lost in a sea of infinite cloud billing. We show you exactly what you use and don't hide any of the costs.

## Getting Started

- Head to [m3o.com](https://m3o.com) and signup for a free account. 
- Generate an API key on the [Settings page](https://m3o.com/settings/keys).
- Browse the APIs on the [Explore page](https://m3o.com/explore).
- Call any API using your token in the `Authorization: Bearer [Token]` header and `https://api.m3o.com/v1/[service]/[endpoint]` url.

## Learn More

- Read the [Announcement](https://blog.m3o.com/2021/06/24/micro-apis-for-everyday-use.html) blog post
- Join the [Slack](https://slack.m3o.com) community
- Join the [Discord](https://discord.gg/TBR9bRjd6Z) channel
- Email [Support](mailto:support@m3o.com) for help

## How it Works

M3O is built on existing public cloud infrastructure with managed kubernetes along with our own [infrastructure automation](https://github.com/m3o/platform) 
and abstraction layer for existing public APIs. We host the open source [Micro](https://github.com/micro/micro) project as our base Cloud OS and use it to 
power all the [Micro Services](https://github.com/micro/services), which provide simpler building blocks for existing cloud primitives.

We then host our own custom dev UX on top and a [backend](https://github.com/m3o/backend) as the management control plane.

## Where Everything Lives

- [m3o.com](https://m3o.com) - a hosted fully managed offering of micro services
- [m3o/m3o](https://github.com/m3o/m3o) - the canonical location for everything on github
- [m3o/cloud](https://github.com/m3o/cloud) - locally hostable dev UX for the website
- [m3o/platform](https://github.com/m3o/platform) - the infrastructure automation for hosted cloud stack
- [m3o/backend](https://github.com/m3o/backend) - the services which power the m3o.com backend
- [micro/micro](https://github.com/micro/micro) - the open source operating system for the cloud
- [micro/services](https://github.com/micro/services) - open source micro services powering m3o.com

## Development

This project is venture funded with a combination of open source services, managed hosting and a team with much distain for AWS.

## Publish APIs

If you'd like to publish your own APIs [fill in this form](https://forms.gle/9SQV6DdLNDzSRQ477) and we'll get back to you.
