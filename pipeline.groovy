def saneBranch = BRANCH.replaceAll("/", "_")
def versionString = VERSION_STRING ? VERSION_STRING : "$TARGET_PROJECT-${saneBranch}-b" + currentBuild.getNumber()

stage("Build") {
    build job: 'at-base-build',
        parameters: [
            booleanParam(name: 'DRYRUN', value: DRYRUN),
            booleanParam(name: 'SKIP_PUBLISHING', value: SKIP_PUBLISHING),
            string(name: 'TARGET_PROJECT', value: TARGET_PROJECT),
            string(name: 'BRANCH', value: BRANCH),
            string(name: 'SCMURL', value: SCMURL),
            string(name: 'VERSION_STRING', value: versionString)
        ]
}
