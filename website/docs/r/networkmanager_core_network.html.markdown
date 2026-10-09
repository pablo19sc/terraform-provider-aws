---
subcategory: "Network Manager"
layout: "aws"
page_title: "AWS: aws_networkmanager_core_network"
description: |-
  Manages a Network Manager Core Network.
---

# Resource: aws_networkmanager_core_network

Manages a Network Manager Core Network.

Use this resource to create and manage a core network within a global network.

## Example Usage

### Basic

Manage the core network policy directly with `policy_document`.

```terraform
resource "aws_networkmanager_global_network" "example" {}

data "aws_networkmanager_core_network_policy_document" "example" {
  core_network_configuration {
    asn_ranges = ["65022-65534"]

    edge_locations {
      location = "us-west-2"
    }
  }

  segments {
    name = "segment"
  }
}

resource "aws_networkmanager_core_network" "example" {
  global_network_id = aws_networkmanager_global_network.example.id
  policy_document   = data.aws_networkmanager_core_network_policy_document.example.json
}
```

### With Base Policy Document

Use `base_policy_document` with `create_base_policy` only when the final policy references attachment IDs or prefix list association names. These references create a Terraform dependency cycle because the attachments and prefix list associations depend on the core network. Apply the final policy after creating those resources with the separate [`aws_networkmanager_core_network_policy_attachment` resource](networkmanager_core_network_policy_attachment.html).

```terraform
data "aws_region" "current" {}

resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "example" {
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.0.0/24"
}

resource "aws_networkmanager_global_network" "example" {}

data "aws_networkmanager_core_network_policy_document" "base" {
  core_network_configuration {
    asn_ranges = ["65022-65534"]

    edge_locations {
      location = data.aws_region.current.region
    }
  }

  segments {
    name = "segment"
  }
}

resource "aws_networkmanager_core_network" "example" {
  global_network_id    = aws_networkmanager_global_network.example.id
  create_base_policy   = true
  base_policy_document = data.aws_networkmanager_core_network_policy_document.base.json
}

resource "aws_networkmanager_vpc_attachment" "example" {
  core_network_id = aws_networkmanager_core_network.example.id
  subnet_arns     = [aws_subnet.example.arn]
  vpc_arn         = aws_vpc.example.arn

  tags = {
    segment = "segment"
  }
}

data "aws_networkmanager_core_network_policy_document" "example" {
  core_network_configuration {
    asn_ranges = ["65022-65534"]

    edge_locations {
      location = data.aws_region.current.region
    }
  }

  segments {
    name = "segment"
  }

  attachment_policies {
    rule_number     = 100
    condition_logic = "or"

    conditions {
      type     = "tag-value"
      operator = "equals"
      key      = "segment"
      value    = "segment"
    }

    action {
      association_method = "constant"
      segment            = "segment"
    }
  }

  segment_actions {
    action                  = "create-route"
    segment                 = "segment"
    destination_cidr_blocks = ["0.0.0.0/0"]
    destinations            = [aws_networkmanager_vpc_attachment.example.id]
  }
}

resource "aws_networkmanager_core_network_policy_attachment" "example" {
  core_network_id = aws_networkmanager_core_network.example.id
  policy_document = data.aws_networkmanager_core_network_policy_document.example.json
}
```

## Argument Reference

The following arguments are required:

* `global_network_id` - (Required) ID of the global network that a core network will be a part of.

The following arguments are optional:

* `base_policy_document` - (Optional, conflicts with `base_policy_regions` and `policy_document`) Base policy document, applied and set to `LIVE` on create. Use with `create_base_policy` when the final policy references attachment IDs or prefix list association names. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
* `base_policy_regions` - (Optional, **Deprecated**, conflicts with `base_policy_document` and `policy_document`) Regions to add to the base policy. Use `base_policy_document` instead. This argument will be removed in the next major version of the provider.
* `create_base_policy` - (Optional, conflicts with `policy_document`) Whether to create and apply a base policy on create or update. Use with `base_policy_document` when the final policy references attachment IDs or prefix list association names. See [With Base Policy Document](#with-base-policy-document).
* `description` - (Optional) Description of the core network.
* `policy_document` - (Optional, conflicts with `base_policy_document`, `base_policy_regions`, and `create_base_policy`) Core network policy document, applied and set to `LIVE`. Preferred way to manage the core network policy. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
* `tags` - (Optional) Key-value tags for the core network. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the core network.
* `arn` - ARN of the core network.
* `created_at` - Timestamp when the core network was created.
* `edges` - Edges created by the `LIVE` core network policy. See [`edges`](#edges) below.
* `segments` - Segments created by the `LIVE` core network policy. See [`segments`](#segments) below.
* `state` - State of the core network.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

### `edges`

`edges` supports the following attributes:

* `asn` - ASN of the core network edge.
* `edge_location` - Region where the core network edge is located.
* `inside_cidr_blocks` - Inside CIDR blocks used by the core network edge.

### `segments`

`segments` supports the following attributes:

* `edge_locations` - Regions containing edges for the segment.
* `name` - Name of the segment.
* `shared_segments` - Names of segments shared with the segment.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `delete` - (Default `30m`)
* `update` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_networkmanager_core_network` using the core network ID. For example:

```terraform
import {
  to = aws_networkmanager_core_network.example
  id = "core-network-0d47f6t230mz46dy4"
}
```

Using `terraform import`, import `aws_networkmanager_core_network` using the core network ID. For example:

```console
% terraform import aws_networkmanager_core_network.example core-network-0d47f6t230mz46dy4
```
