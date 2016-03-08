SELECT B.BLOCK_ID, B.BLOCK_NAME, B.OPEN_FOR_WRITING, 
        B.BLOCK_SIZE, B.FILE_COUNT,
        B.DATASET_ID, DS.DATASET,
        B.ORIGIN_SITE_NAME, B.CREATION_DATE, B.CREATE_BY,
        B.LAST_MODIFICATION_DATE, B.LAST_MODIFIED_BY
FROM {{.Owner}}.BLOCKS B JOIN {{.Owner}}.DATASETS DS ON DS.DATASET_ID = B.DATASET_ID 