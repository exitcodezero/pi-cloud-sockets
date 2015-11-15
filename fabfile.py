from fabric.api import env, run, local, sudo, settings


def build_local():
    local('docker-compose run app go build -v')


def copy_app():
    local('scp picloud {0}@{1}:/home/{0}'.format(env.user, env.hosts[0]))


def stop_service():
    with settings(warn_only=True):
        sudo('service pi-cloud stop')


def remove_old_app():
    run('rm pi-cloud')


def rename_new_app():
    run('mv picloud pi-cloud')


def start_service():
    sudo('service pi-cloud start')


def remove_local():
    local('rm picloud')


def deploy():
    copy_app()
    stop_service()
    remove_old_app()
    rename_new_app()
    start_service()


def build_deploy():
    build_local()
    deploy()
    remove_local()
