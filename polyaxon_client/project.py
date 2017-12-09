# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

import sys

from polyaxon_schemas.experiment import ExperimentConfig
from polyaxon_schemas.project import ProjectConfig, PolyaxonSpecConfig
from polyaxon_schemas.polyaxonfile.logger import logger

from polyaxon_client.base import PolyaxonClient
from polyaxon_client.exceptions import PolyaxonException, AuthenticationError, NotFoundError


class ProjectClient(PolyaxonClient):
    """Client to get projects from the server"""
    ENDPOINT = "/projects"

    def list_projects(self, page=1):
        try:
            response = self.get(self._get_url(), params=self.get_page(page=page))
            projects_dict = response.json()
            return [ProjectConfig.from_dict(project) 
                    for project in projects_dict.get("results", [])]
        except PolyaxonException as e:
            logger.info("Error while retrieving projects: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return []

    def get_project(self, project_uuid):
        request_url = self._build_url(self._get_url(), project_uuid)
        try:
            response = self.get(request_url)
            return ProjectConfig.from_dict(response.json())
        except NotFoundError:
            return None

    def create_project(self, project_config):
        try:
            response = self.post(self._get_url(), json=project_config.to_dict())
            return ProjectConfig.from_dict(response.json())
        except PolyaxonException as e:
            logger.info("Error while creating project: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return None

    def update_project(self, project_uuid, patch_dict):
        request_url = self._build_url(self._get_url(), project_uuid)
        try:
            response = self.patch(request_url, json=patch_dict)
            return ProjectConfig.from_dict(response.json())
        except PolyaxonException as e:
            logger.info("Error while updating project: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return None

    def delete_project(self, project_uuid):
        request_url = self._build_url(self._get_url(), project_uuid)
        try:
            response = self.delete(request_url)
            return response
        except PolyaxonException as e:
            logger.info("Error while updating project: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return None

    def upload_repo(self, project_uuid, files, files_size=None):
        """Uploads code data related for this project from the current dir."""
        request_url = self._build_url(self._get_url(), project_uuid, 'repo', 'upload')

        try:
            response = self.upload(request_url, files=files, files_size=files_size)
            return response
        except PolyaxonException as e:
            logger.info("Error while updating project: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return None

    def list_specs(self, project_uuid, page=1):
        """Fetch list of specs related to this project."""
        request_url = self._build_url(self._get_url(), project_uuid, 'specs')

        try:
            response = self.get(request_url, params=self.get_page(page=page))
            projects_dict = response.json()
            return [PolyaxonSpecConfig.from_dict(spec)
                    for spec in projects_dict.get("results", [])]
        except PolyaxonException as e:
            logger.info("Error while retrieving projects: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return []

    def create_spec(self, project_uuid, spec_config):
        request_url = self._build_url(self._get_url(), project_uuid, 'specs')

        try:
            response = self.post(request_url, json=spec_config.to_dict())
            return PolyaxonSpecConfig.from_dict(response.json())
        except PolyaxonException as e:
            logger.info("Error while creating project: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return None

    def update_spec(self, project_uuid, spec_uuid, patch_dict):
        request_url = self._build_url(self._get_url(), project_uuid, 'specs', spec_uuid)

        try:
            response = self.patch(request_url, json=patch_dict)
            return PolyaxonSpecConfig.from_dict(response.json())
        except PolyaxonException as e:
            logger.info("Error while updating project: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return None

    def delete_spec(self, project_uuid, spec_uuid):
        request_url = self._build_url(self._get_url(), project_uuid, 'specs', spec_uuid)
        try:
            response = self.delete(request_url)
            return response
        except PolyaxonException as e:
            logger.info("Error while updating project: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return None

    def list_experiments(self, uuid, page=1):
        """Fetch list of specs related to this project."""
        request_url = self._build_url(self._get_url(), uuid, 'experiments')

        try:
            response = self.get(request_url, params=self.get_page(page=page))
            projects_dict = response.json()
            return [ExperimentConfig.from_dict(spec)
                    for spec in projects_dict.get("results", [])]
        except PolyaxonException as e:
            logger.info("Error while retrieving projects: {}".format(e.message))
            if isinstance(e, AuthenticationError):
                # exit now since there is nothing we can do without login
                sys.exit(1)
            return []
