# Update CAPI Providers with CAPI beta releases

`provider_issues` is a go utility intended to open git issues on provider repos about the latest CAPI beta release.

## Pre-requisites

- Create a github token with the below access and save it in your environment as `GITHUB_ISSUE_OPENER_TOKEN`.
  - `repo:status` - Grants access to commit status on public and private repositories.
  - `repo_deployment` - Grants access to deployment statuses on public and private repositories.
  - `public_repo` - Grants access to public repositories
- Create `PROVIDER_ISSUES_DRY_RUN` environment variable and set it to `"true"` run the utility in dry run mode. Set it to `"false"` to create issues on the provider repositories.

  ```sh
  export PROVIDER_ISSUES_DRY_RUN="false"
  ```

- Decide upon the title of the issue to be opened. Update it in the utility.
- Update the issue body. Make sure to include following release and upcoming schedule.
  
  ```md
  <!-- body -->
  <!-- TODO: remove all TODOs before running this utility -->
  <!-- TODO: update CAPI release semver -->
  CAPI v1.x.0-beta.0 has been released and is ready for testing.
  Looking forward to your feedback before CAPI 1.x.0 release!

  ## For quick reference

  <!-- body -->
  <!-- TODO: CAPI release notes -->
  - [CAPI v1.x.0-beta.0 release notes](https://github.com/kubernetes-sigs/cluster-api/releases/tag/v1.x.0-beta.0)
  - [Shortcut to CAPI git issues](https://github.com/kubernetes-sigs/cluster-api/issues)

  ## Following are the planned dates for the upcoming releases

  <!-- TODO: update CAPI release timeline -->
  |Release|Expected Date|
  |-----|-----|
  |v1.5.0-beta.x | Tuesday 5th July 2023|
  |release-1.5 branch created (Begin [Code Freeze])|Tuesday 11th July 2023|
  |v1.5.0-rc.0 released|Tuesday 11th July 2023|
  | release-1.5 jobs created | Tuesday 11th July 2023 |
  | v1.5.0-rc.x released | Tuesday 18th July 2023 |
  | v1.5.0 released | Tuesday 25th July 2023 |
  <!-- body -->

  <!-- [List of CAPI providers](https://github.com/kubernetes-sigs/cluster-api/blob/main/docs/release/release-tasks.md#communicate-beta-to-providers) -->
  ```

## How to run the tool

- From the root of the project Cluster API, run below to run the utility.

  ```sh
  make provider-issues
  ```

  - Note: the `provider_issues.go` utility will always run in dry-run mode **unless** the env variable `PROVIDER_ISSUES_DRY_RUN` is set to `"false"`.
  - When **running in dry run mode**, you will see below output.

    ```bash
    ###############################################
    This script will run in dry run mode.
    To run it for real, set the PROVIDER_ISSUES_DRY_RUN environment variable to "false".
    ###############################################

    Issues will be created for the following repositories:
    - kubernetes-sigs/cluster-api-addon-provider-helm
    - kubernetes-sigs/cluster-api-provider-aws
    - kubernetes-sigs/cluster-api-provider-azure
    - kubernetes-sigs/cluster-api-provider-cloudstack
    - kubernetes-sigs/cluster-api-provider-digitalocean
    - kubernetes-sigs/cluster-api-provider-gcp
    - kubernetes-sigs/cluster-api-provider-kubemark
    - kubernetes-sigs/cluster-api-provider-kubevirt
    - kubernetes-sigs/cluster-api-provider-ibmcloud
    - kubernetes-sigs/cluster-api-provider-nested
    - oracle/cluster-api-provider-oci
    - kubernetes-sigs/cluster-api-provider-openstack
    - kubernetes-sigs/cluster-api-operator
    - kubernetes-sigs/cluster-api-provider-packet
    - kubernetes-sigs/cluster-api-provider-vsphere

    With the following issue body:

    <!-- body -->
    <!-- TODO: remove all TODOs before running this utility -->
    <!-- TODO: update CAPI release semver -->
    CAPI v1.x.0-beta.0 has been released and is ready for testing.
    Looking forward to your feedback before CAPI 1.x.0 release!

    ## For quick reference

    <!-- body -->
    <!-- TODO: CAPI release notes -->
    - [CAPI v1.x.0-beta.0 release notes](https://github.com/kubernetes-sigs/cluster-api/releases/tag/v1.x.0-beta.0)
    - [Shortcut to CAPI git issues](https://github.com/kubernetes-sigs/cluster-api/issues)

    ## Following are the planned dates for the upcoming releases

    <!-- TODO: update CAPI release timeline -->
    |Release|Expected Date|
    |-----|-----|
    |v1.5.0-beta.x | Tuesday 5th July 2023|
    |release-1.5 branch created (Begin [Code Freeze])|Tuesday 11th July 2023|
    |v1.5.0-rc.0 released|Tuesday 11th July 2023|
    | release-1.5 jobs created | Tuesday 11th July 2023 |
    | v1.5.0-rc.x released | Tuesday 18th July 2023 |
    | v1.5.0 released | Tuesday 25th July 2023 |

    <!-- [List of CAPI providers](https://github.com/kubernetes-sigs/cluster-api/blob/main/docs/release/release-tasks.md#communicate-beta-to-providers) -->
    <!-- body -->


    DRY RUN: issue(s) body will not be posted.
    Exiting...
    ```
  
  - When running with `PROVIDER_ISSUES_DRY_RUN` set to `"false"`
  
    ```bash
    Issues will be created for the following repositories:
    - kubernetes-sigs/cluster-api-addon-provider-helm
    - kubernetes-sigs/cluster-api-provider-aws
    - kubernetes-sigs/cluster-api-provider-azure
    - kubernetes-sigs/cluster-api-provider-cloudstack
    - kubernetes-sigs/cluster-api-provider-digitalocean
    - kubernetes-sigs/cluster-api-provider-gcp
    - kubernetes-sigs/cluster-api-provider-kubemark
    - kubernetes-sigs/cluster-api-provider-kubevirt
    - kubernetes-sigs/cluster-api-provider-ibmcloud
    - kubernetes-sigs/cluster-api-provider-nested
    - oracle/cluster-api-provider-oci
    - kubernetes-sigs/cluster-api-provider-openstack
    - kubernetes-sigs/cluster-api-operator
    - kubernetes-sigs/cluster-api-provider-packet
    - kubernetes-sigs/cluster-api-provider-vsphere

    With the following issue body:

    <!-- body -->
    <!-- TODO: remove all TODOs before running this utility -->
    <!-- TODO: update CAPI release semver -->
    CAPI v1.x.0-beta.0 has been released and is ready for testing.
    Looking forward to your feedback before CAPI 1.x.0 release!

    ## For quick reference

    <!-- body -->
    <!-- TODO: CAPI release notes -->
    - [CAPI v1.x.0-beta.0 release notes](https://github.com/kubernetes-sigs/cluster-api/releases/tag/v1.x.0-beta.0)
    - [Shortcut to CAPI git issues](https://github.com/kubernetes-sigs/cluster-api/issues)

    ## Following are the planned dates for the upcoming releases

    <!-- TODO: update CAPI release timeline -->
    |Release|Expected Date|
    |-----|-----|
    |v1.5.0-beta.x | Tuesday 5th July 2023|
    |release-1.5 branch created (Begin [Code Freeze])|Tuesday 11th July 2023|
    |v1.5.0-rc.0 released|Tuesday 11th July 2023|
    | release-1.5 jobs created | Tuesday 11th July 2023 |
    | v1.5.0-rc.x released | Tuesday 18th July 2023 |
    | v1.5.0 released | Tuesday 25th July 2023 |

    <!-- [List of CAPI providers](https://github.com/kubernetes-sigs/cluster-api/blob/main/docs/release/release-tasks.md#communicate-beta-to-providers) -->
    <!-- body -->

    Continue? (y/n)
    y

    Issue created for repository 'kubernetes-sigs/cluster-api-addon-provider-helm'
    URL: https://github.com/kubernetes-sigs/cluster-api-addon-provider-helm/issues/1234

    Issue created for repository 'kubernetes-sigs/cluster-api-provider-aws'
    URL: https://github.com/kubernetes-sigs/cluster-api-provider-aws/issues/1234

    Issue created for repository 'kubernetes-sigs/cluster-api-provider-azure'
    URL: https://github.com/kubernetes-sigs/cluster-api-provider-azure/issues/1234
    ```
