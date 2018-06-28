from django.conf import settings

from libs.paths.exceptions import VolumeNotFoundError


def get_outputs_paths(persistence_outputs):
    # If no persistence is defined we mount the first one as default
    persistence_outputs = persistence_outputs or list(settings.PERSISTENCE_OUTPUTS.keys())[0]
    if persistence_outputs not in settings.PERSISTENCE_OUTPUTS:
        raise VolumeNotFoundError('Outputs volume with name `{}` was defined in specification, '
                                  'but was not found'.format(persistence_outputs))
    if 'mountPath' not in settings.PERSISTENCE_OUTPUTS[persistence_outputs]:
        raise VolumeNotFoundError('Outputs volume with name `{}` '
                                  'does not define a mountPath.'.format(persistence_outputs))

    return settings.PERSISTENCE_OUTPUTS[persistence_outputs]['mountPath']
