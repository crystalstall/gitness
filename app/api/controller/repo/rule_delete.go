// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repo

import (
	"context"
	"fmt"

	"github.com/harness/gitness/app/auth"
	"github.com/harness/gitness/types/enum"
)

// RuleDelete deletes a protection rule by UID.
func (c *Controller) RuleDelete(ctx context.Context,
	session *auth.Session,
	repoRef string,
	uid string,
) error {
	repo, err := c.getRepoCheckAccess(ctx, session, repoRef, enum.PermissionRepoEdit, false)
	if err != nil {
		return err
	}

	r, err := c.ruleStore.FindByUID(ctx, nil, &repo.ID, uid)
	if err != nil {
		return fmt.Errorf("failed to find repository-level protection rule by uid: %w", err)
	}

	err = c.ruleStore.Delete(ctx, r.ID)
	if err != nil {
		return fmt.Errorf("failed to delete repository-level protection rule: %w", err)
	}

	return nil
}
