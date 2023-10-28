package dao

import (
	"github.com/sirupsen/logrus"
	"scan_project/internal/model"
)

type ScansDao struct {
	jobsScans     map[daoKey][]model.JobScan
	servicesScans map[daoKey][]model.ServiceScan
	logger        *logrus.Entry
}

func NewScansDao(logger *logrus.Entry) ScansDao {
	return ScansDao{
		jobsScans:     make(map[daoKey][]model.JobScan),
		servicesScans: make(map[daoKey][]model.ServiceScan),
		logger:        logger,
	}
}

type daoKey struct {
	kubeconfigName string
	namespace      string
}

func (sd *ScansDao) GetJobsScans(configName string, namespace string) []model.JobScan {
	scans := sd.jobsScans[daoKey{
		kubeconfigName: configName,
		namespace:      namespace,
	}]
	if scans != nil {
		return scans
	} else {
		return make([]model.JobScan, 0)
	}
}

func (sd *ScansDao) UpdateJobsScans(configName string, namespace string, jobsScans []model.JobScan) error {
	key := daoKey{
		kubeconfigName: configName,
		namespace:      namespace,
	}
	sd.jobsScans[key] = jobsScans
	return nil
}

func (sd *ScansDao) GetServicesScans(configName string, namespace string) []model.ServiceScan {
	scans := sd.servicesScans[daoKey{
		kubeconfigName: configName,
		namespace:      namespace,
	}]
	if scans != nil {
		return scans
	} else {
		return make([]model.ServiceScan, 0)
	}
}

func (sd *ScansDao) UpdateServicesScans(configName string, namespace string, servicesScans []model.ServiceScan) error {
	key := daoKey{
		kubeconfigName: configName,
		namespace:      namespace,
	}
	sd.servicesScans[key] = servicesScans
	return nil
}
