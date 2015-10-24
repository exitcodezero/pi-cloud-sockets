import os
from fabric.api import env, run, local, sudo, settings
from fabric.contrib.console import confirm


DEPLOY_HOST = os.getenv('DEPLOY_HOST')
assert DEPLOY_HOST

DEPLOY_USER = os.getenv('DEPLOY_USER')
assert DEPLOY_USER


env.hosts = [DEPLOY_HOST]
env.user = DEPLOY_USER


def copy_app():
    cmd = 'scp app {0}@{1}:/home/{0}'.format(DEPLOY_USER, DEPLOY_HOST)
    local(cmd)


def stop_service():
    with settings(warn_only=True):
        sudo('service pi-cloud stop')


def remove_old_app():
    run('rm pi-cloud')


def rename_new_app():
    run('mv app pi-cloud')


def start_service():
    sudo('service pi-cloud start')


def deploy():
    copy_app()
    stop_service()
    remove_old_app()
    rename_new_app()
    start_service()
