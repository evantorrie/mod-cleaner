const core = require('@actions/core');
const github = require('@actions/github');
const cli = require('@actions/cli');
const glob = require('@actions/glob');

try {
  // `files` input defined in action metadata file
    const files = core
          .getInput('files')
          .split("\n")
          .filter(x => x !== "");
    
    console.log(`files input: ${files}!`);
    const allFiles = findFiles(files);
    const time = (new Date()).toTimeString();
    core.setOutput("changedfiles", time);
    // Get the JSON webhook payload for the event that triggered the workflow
    const payload = JSON.stringify(github.context.payload, undefined, 2);
    console.log(`\nThe event payload: ${payload}`);
} catch (error) {
  core.setFailed(error.message);
}

async function findFiles(patterns) {
    for (const gomod of files) {
        console.log(`gomod-pattern: ${gomod}`);
        const globber = await glob.create(gomod);
        const files = await globber.glob();
        console.log(`expanded = ${files}`);
    }
}

