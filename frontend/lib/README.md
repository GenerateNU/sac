# Generate SAC Frontend Library:

This repository holds the common types, apis, and components for frontend library used by the Student Activity Calendar.

To install the page, navigate to the target repo and initiate the following command:

-   `yarn add @generatesac/lib`

## Development / Testing

When contributing to the package, you will likely want to test your code locally before pushing to remote. To link it with real-time updates, following the steps below:

1. Navigate to `/frontend/lib` and run `yarn link`
2. Mavigate to the target repo, ppen a second terminal and run `yarn link "@generatesac/lib"`
3. Navigate back to `/frontend/lib` and run `yarn watch`.

When you are ready to finish local development:

1. Exit out of "watch mode" using `cmd + C` / `ctrl + C`
1. Navigate to `/frontend/lib` and run `yarn unlink`
1. Navigate to the target repo and run `yarn link "@generatesac/lib"`

## Contributing

When you are ready to push new code to the library, make sure to bump the version in the `package.json`. We are using semantic versioning. If you forget to do this, the `frontend/lib` GitHub Action will fail on push.

Once the code is merged into `main`, a new version of the package will automatically be built and deployed. Monitor GitHub Actions to ensure a proper release. In the event that the workflow fails, you can manually release the package using the commands below:

1. `yarn build`
2. `yarn publish`
